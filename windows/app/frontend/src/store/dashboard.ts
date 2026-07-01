import { create } from 'zustand';
import { SystemInfoUpdate, SecurityInfoUpdate, AppInfoUpdate } from '../../wailsjs/go/main/App';

export interface DeviceInfoItem {
  title: string;
  value: string;
}

export interface SecurityStatusItem {
  title: string;
  value: boolean;
}

export interface AppInfoItem {
  title: string;
  value: string;
}

// Backend Response Key Mirrors
interface BackendSecurityInfo {
  windowsDefender: boolean;
  firewall: boolean;
  tpm: boolean;
  bitLocker: boolean;
}

interface BackendSystemInfo {
  hostname: string;
  domain: string;
  user: string;
  os: string;
}

interface BackendAppInfo {
  id: string;
  deviceToken: string;
  name: string;
  version: string;
  build: string;
  tag: string;
  buildType: string;
  company: string;
  website: string;
  endpoint: string;
  description: string;
}

const SECURITY_MAP: { key: keyof BackendSecurityInfo; title: string }[] = [
  { key: 'windowsDefender', title: 'Windows Defender' },
  { key: 'firewall', title: 'Firewall' },
  { key: 'tpm', title: 'TPM' },
  { key: 'bitLocker', title: 'BitLocker' },
];

const DEVICE_MAP: { key: keyof BackendSystemInfo; title: string }[] = [
  { key: 'hostname', title: 'Hostname' },
  { key: 'domain', title: 'Domain' },
  { key: 'user', title: 'User' },
  { key: 'os', title: 'OS' },
];

interface SecurityInfoState {
  securityStatus: SecurityStatusItem[];
  securityLoading: boolean;
  error: unknown;
  fetchSecurityData: () => Promise<void>;
}

export const useSecurityInfoStore = create<SecurityInfoState>((set) => ({
  securityStatus: [],
  securityLoading: true,
  error: null,
  fetchSecurityData: async () => {
    set({ securityLoading: true, error: null });
    try {
      const securityInfo: BackendSecurityInfo = await SecurityInfoUpdate();
      set({
        securityStatus: SECURITY_MAP.map(({ key, title }) => ({
          title,
          value: Boolean(securityInfo[key]),
        })),
        securityLoading: false,
      });
    } catch (error) {
      set({ error, securityLoading: false });
    }
  },
}));

interface DeviceInfoState {
  deviceInfo: DeviceInfoItem[];
  deviceLoading: boolean;
  error: unknown;
  fetchData: () => Promise<void>;
}

export const useDeviceInfoStore = create<DeviceInfoState>((set) => ({
  deviceInfo: [],
  deviceLoading: true,
  error: null,
  fetchData: async () => {
    set({ deviceLoading: true, error: null });
    try {
      const systemInfo: BackendSystemInfo = await SystemInfoUpdate();
      set({
        deviceInfo: DEVICE_MAP.map(({ key, title }) => ({
          title,
          value: String(systemInfo[key]),
        })),
        deviceLoading: false,
      });
    } catch (error) {
      set({ error, deviceLoading: false }); // FIXED: Corrected error target loader variable context
    }
  },
}));



interface AppInfoState {
  appInfoLoading: boolean;
  error: unknown;
  appInfo: BackendAppInfo;  
  fetchData: () => Promise<void>;
}

export const useAppInfoStore = create<AppInfoState>((set) => ({
  appInfoLoading: true,
  error: null,
  appInfo: {} as BackendAppInfo,
  fetchData: async () => {
    set({ appInfoLoading: true, error: null });
    try {
      const appDetails: BackendAppInfo = await AppInfoUpdate();
      set({
        appInfo: appDetails,
        appInfoLoading: false,
      });
    } catch (error) {
      set({ error, appInfoLoading: false });
    }
  },
}));