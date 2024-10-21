'use client';

import React, {
  createContext, ReactNode, useContext, useMemo, useState,
} from 'react';
import { LoadedProfile } from '../services/dtos/responses/default-response';

interface AuthContextType {
  profile: LoadedProfile | null;
  setProfile: (loadedProfile: LoadedProfile) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [profile, setProfile] = useState<LoadedProfile | null>(null);

  const saveProfile = (loadedProfile: LoadedProfile) => {
    setProfile(loadedProfile);
  };

  const logout = () => {
    setProfile(null);
  };

  const contextValue = useMemo(() => ({
    profile,
    setProfile: saveProfile,
    logout,
  }), [profile]);

  return (
    <AuthContext.Provider value={contextValue}>
      {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
