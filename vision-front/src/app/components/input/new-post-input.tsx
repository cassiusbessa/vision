'use client';

import React from 'react';
import { useAuth } from '@/app/state/auth-context';
import UserIcon from '../user-icon';

export default function NewPostInput({ className = '' }:{ className: string }) {
  const { me } = useAuth();
  const handleOpenModal = () => {
    const element = document.getElementById('create_projet_modal') as any;
    if (element) element.showModal();
  };

  const userImage = me?.profile?.image || 'https://e7.pngegg.com/pngimages/84/165/png-clipart-united-states-avatar-organization-information-user-avatar-service-computer-wallpaper-thumbnail.png';

  return (
    <div className={`w-3/4 md:w-8/12 lg:w-1/2 mt-6 flex flex-row gap-2 bg-base-100 items-center p-4 ${className}`}>
      <UserIcon src={userImage} />
      <label htmlFor="post-project" className="input input-bordered flex items-center gap-2 w-full bg-base-100 border-1 border-solid rounded-2xl border-slate-300">
        <span className="text-xl sr-only">🚀</span>
        <input id="post-project" type="text" className="grow input-bordered" placeholder="Compartilhe seu projeto!" onClick={handleOpenModal} />
      </label>
    </div>
  );
}
