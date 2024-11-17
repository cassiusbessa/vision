'use client';

/* eslint-disable react/jsx-props-no-spreading */

import { loadMe } from '@/app/services/account';

import { getToken } from '@/app/services/token';
import { useAuth } from '@/app/state/auth-context';
import React, { useEffect, ComponentType } from 'react';

function withProfile<T>(WrappedComponent: ComponentType<T>) {
  return function WithProfile(props: T & JSX.IntrinsicAttributes) {
    const { setMe, me } = useAuth();

    useEffect(() => {
      const token = getToken();

      if (token && !me) {
        loadMe(token).then((response) => {
          if (response.ok && response.data) {
            setMe(response.data);
          }
          return me;
        });
      }
    }, [setMe, me]);
    return <WrappedComponent {...props} />;
  };
}

export default withProfile;
