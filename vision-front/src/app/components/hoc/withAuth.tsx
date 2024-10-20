/* eslint-disable react/jsx-props-no-spreading */
import { getToken } from '@/app/services/token';
import { useRouter } from 'next/navigation';
import React, { useEffect, ComponentType } from 'react';

function withAuth<T>(WrappedComponent: ComponentType<T>) {
  return function WithAuth(props: T & JSX.IntrinsicAttributes) {
    const router = useRouter();

    useEffect(() => {
      const token = getToken();

      if (!token) {
        router.push('/login');
      }
    }, [router]);
    return <WrappedComponent {...props} />;
  };
}

export default withAuth;
