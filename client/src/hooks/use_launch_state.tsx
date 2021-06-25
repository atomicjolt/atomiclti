import { useState } from 'react';
import jwtDecode from 'jwt-decode';

export interface LaunchState {
  contextId: string;
}

const LAUNCH_STATE: LaunchState = jwtDecode(window.LAUNCH_TOKEN);

export default (): LaunchState => {
  const [launchState] = useState<LaunchState>(LAUNCH_STATE);

  return launchState;
};
