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

import React, { PropsWithChildren, useEffect } from 'react';
import Logger from 'shared/libs/logger';
import useAttempt from 'shared/hooks/useAttemptNext';
import { getErrMessage } from 'shared/utils/errorType';
import { Box, Indicator } from 'design';

import session from 'teleport/services/websession';
import { ApiError } from 'teleport/services/api/parseError';
import { StyledIndicator } from 'teleport/Main';

import { ErrorDialog } from './ErrorDialogue';

const logger = Logger.create('/components/Authenticated');
const Authenticated: React.FC<PropsWithChildren> = ({ children }) => {
  const { attempt, setAttempt } = useAttempt('processing');

  useEffect(() => {
    const checkIfUserIsAuthenticated = async () => {
      if (!session.isValid()) {
        logger.warn('invalid session');
        session.logoutWithoutSlo(true /* rememberLocation */);
        return;
      }

      try {
        const result = await session.validateCookieAndSession();
        if (result.hasDeviceExtensions) {
          session.setIsDeviceTrusted();
        }
        setAttempt({ status: 'success' });
      } catch (e) {
        if (e instanceof ApiError && e.response?.status == 403) {
          logger.warn('invalid session');
          session.logoutWithoutSlo(true /* rememberLocation */);
          // No need to update attempt, as `logout` will
          // redirect user to login page.
          return;
        }
        // Error unrelated to authentication failure (network blip).
        setAttempt({ status: 'failed', statusText: getErrMessage(e) });
      }
    };

    checkIfUserIsAuthenticated();
  }, []);

  if (attempt.status === 'success') {
    return <>{children}</>;
  }

  if (attempt.status === 'failed') {
    return <ErrorDialog errMsg={attempt.statusText} />;
  }

  return (
    <Box textAlign="center">
      <StyledIndicator>
        <Indicator />
      </StyledIndicator>
    </Box>
  );
};

export default Authenticated;
