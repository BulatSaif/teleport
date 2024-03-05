/**
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
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

import { FileTransferListeners } from 'shared/components/FileTransfer';

import { FileTransferDirection } from 'gen-proto-ts/teleport/lib/teleterm/v1/service_pb';

import { FileTransferRequest, TshdClient } from 'teleterm/services/tshd/types';
import { UsageService } from 'teleterm/ui/services/usage';
import { objectifyAbortSignal } from 'teleterm/services/tshd/grpcContextBridgeClient';

export class FileTransferService {
  constructor(
    private tshClient: TshdClient,
    private usageService: UsageService
  ) {}

  transferFile(
    options: FileTransferRequest,
    abortController: AbortController
  ): FileTransferListeners {
    const listeners = this.tshClient.transferFile(
      options,
      objectifyAbortSignal(abortController.signal)
    );
    if (options.direction === FileTransferDirection.DOWNLOAD) {
      this.usageService.captureFileTransferRun(options.serverUri, {
        isUpload: false,
      });
    }
    if (options.direction === FileTransferDirection.UPLOAD) {
      this.usageService.captureFileTransferRun(options.serverUri, {
        isUpload: true,
      });
    }
    return listeners;
  }
}
