/*
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package tbot

import (
	"context"
	"log/slog"
	"math"
	"time"

	"github.com/gravitational/trace"

	"github.com/gravitational/teleport/api/utils/retryutils"
)

type runOnIntervalConfig struct {
	name                 string
	f                    func(ctx context.Context) error
	reloadCh             chan struct{}
	log                  *slog.Logger
	interval             time.Duration
	retryLimit           int
	exitOnRetryExhausted bool
	waitBeforeFirstRun   bool
}

// runOnInterval runs a function on a given interval, with retries and jitter.
//
// TODO(noah): Emit Prometheus metrics for:
// - Success/Failure of attempts
// - Time taken to execute attempt
// - Time of next attempt
func runOnInterval(ctx context.Context, cfg runOnIntervalConfig) error {
	ticker := time.NewTicker(cfg.interval)
	jitter := retryutils.NewJitter()

	// If no reload channel is provided, create one that never yields a value.
	if cfg.reloadCh == nil {
		cfg.reloadCh = make(chan struct{})
	}
	log := cfg.log.With("task", cfg.name)

	firstRun := true

	defer ticker.Stop()
	for {
		if !firstRun || (firstRun && cfg.waitBeforeFirstRun) {
			select {
			case <-ctx.Done():
				return nil
			case <-ticker.C:
			case <-cfg.reloadCh:
			}
		}
		firstRun = false

		var err error
		for attempt := 1; attempt <= cfg.retryLimit; attempt++ {
			log.InfoContext(
				ctx,
				"Attempting task",
				"attempt", attempt,
				"retry_limit", cfg.retryLimit,
			)
			err = cfg.f(ctx)
			if err == nil {
				break
			}

			if attempt != cfg.retryLimit {
				// exponentially back off with jitter, starting at 1 second.
				backoffTime := time.Second * time.Duration(math.Pow(2, float64(attempt-1)))
				backoffTime = jitter(backoffTime)
				cfg.log.WarnContext(
					ctx,
					"Task failed. Backing off and retrying",
					"attempt", attempt,
					"retry_limit", cfg.retryLimit,
					"backoff", backoffTime,
					"error", err,
				)
				select {
				case <-ctx.Done():
					return nil
				case <-time.After(backoffTime):
				}
			}
		}
		if err != nil {
			if cfg.exitOnRetryExhausted {
				log.ErrorContext(
					ctx,
					"All retry attempts exhausted. Exiting",
					"error", err,
					"retry_limit", cfg.retryLimit,
				)
				return trace.Wrap(err)
			}
			log.WarnContext(
				ctx,
				"All retry attempts exhausted. Will wait for next interval",
				"retry_limit", cfg.retryLimit,
				"interval", cfg.interval,
			)
		} else {
			log.InfoContext(
				ctx,
				"Task succeeded. Waiting interval",
				"interval", cfg.interval,
			)
		}
	}
}
