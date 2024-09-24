'use client';

import React from 'react';
import ProfileForm from './profile-form';

function ProfileManager() {
  return (
    <div className="my-4 w-full flex justify-center">
      <div className="w-10/12 max-w-[560px] h-fit bg-[#3c3c3c] border border-slate-400 rounded-3xl py-8 px-4 md:px-8 md:py-12 shadow-lg backdrop-filter backdrop-blur-sm relative">
        <div className="text-center mb-6 text-4x1 text-xl font-medium">
          <h1>Crie o seu Perfil</h1>
        </div>
        <ProfileForm />
      </div>
    </div>
  );
}

export default ProfileManager;
