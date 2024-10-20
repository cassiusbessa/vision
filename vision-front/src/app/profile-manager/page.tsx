'use client';

import React from 'react';
import { Josefin_Sans } from 'next/font/google';
import Header from '../components/header';
import ProfileManager from '../components/profile-manager/profile-manager';
import withAuth from '../components/hoc/withAuth';

const inter = Josefin_Sans({ subsets: ['latin'] });

function ProfileManagerPage() {
  return (
    <div className={`${inter.className} bg-cover bg-gradient-to-b from-[#291D32] via-[#392039] to-[#291D32] flex flex-col items-center min-h-screen`}>
      <Header />
      <ProfileManager />
    </div>
  );
}

export default withAuth(ProfileManagerPage);
