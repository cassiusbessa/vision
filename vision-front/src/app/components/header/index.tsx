'use client';

import React from 'react';
import { useAuth } from '@/app/state/auth-context';
import MobileHeader from './mobile-header';
import ProfileIconDropDown from './profile-icon-dropdown';

export default function Header() {
  const { profile, logout } = useAuth();

  const profileImage = profile?.profile.image || 'https://e7.pngegg.com/pngimages/84/165/png-clipart-united-states-avatar-organization-information-user-avatar-service-computer-wallpaper-thumbnail.png';

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
