'use client';

import React from 'react';
import { useAuth } from '@/app/state/auth-context';
import MobileHeader from './mobile-header';
import ProfileIconDropDown from './profile-icon-dropdown';

export default function Header() {
  const { profile, logout } = useAuth();

  const profileImage = profile?.profile.image || 'https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg';
  console.log(profile);

  return (
    <div className="navbar bg-cover bg-gradient-to-b from-[#291D32] via-[#392039] to-[#291D32]">
      <div className="navbar-start">
        <MobileHeader />
        <a className="btn btn-ghost text-xl" href="/">daisyUI</a>
      </div>
      <div className="navbar-end">
        <ProfileIconDropDown
          profileImage={profileImage}
          profileLink={profile?.profile.link || ''}
          logout={logout}
        />
      </div>
    </div>
  );
}
