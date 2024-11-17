'use client';

import React, {
  createContext, ReactNode, useContext, useMemo, useState,
} from 'react';
import { Me } from '../services/dtos/responses/default-response';
import { removeTokenLocalStorage, removeTokenSessionStorage } from '../services/token';

interface AuthContextType {
  me: Me | null;
  setMe: (me: Me) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [me, setMe] = useState<Me | null>(null);

  const saveMe = (meResponse: Me) => {
    setMe(meResponse);
  };

  const logout = () => {
    setMe(null);
    removeTokenLocalStorage();
    removeTokenSessionStorage();
  };

  const contextValue = useMemo(() => ({
    me,
    setMe: saveMe,
    logout,
  }), [me]);

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
