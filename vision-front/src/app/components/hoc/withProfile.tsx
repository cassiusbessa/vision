'use client';

/* eslint-disable react/jsx-props-no-spreading */

import { loadProfileByToken } from '@/app/services/profile';
import { getToken } from '@/app/services/token';
import { useAuth } from '@/app/state/auth-context';
import React, { useEffect, ComponentType } from 'react';

function withProfile<T>(WrappedComponent: ComponentType<T>) {
  return function WithProfile(props: T & JSX.IntrinsicAttributes) {
    const { setProfile, profile } = useAuth();

    useEffect(() => {
      const token = getToken();

      if (token && !profile) {
        loadProfileByToken().then((response) => {
          if (response.ok && response.data) {
            setProfile(response.data);
          }
          return profile;
        });
      }
    }, [setProfile, profile]);
    return <WrappedComponent {...props} />;
  };
}

export default withProfile;
