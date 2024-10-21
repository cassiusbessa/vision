import React from 'react';
import { useRouter } from 'next/navigation';
import UserIcon from '../user-icon';

export default function ProfileIconDropDown({ profileImage, profileLink, logout }:
{ profileImage: string, profileLink: string, logout: () => void, }) {
  const router = useRouter();
  const handleLogout = () => {
    logout();
    router.push('/login');
  };

  const redirectProfile = () => {
    router.push(`/profile/${profileLink}`);
  };

  return (
    <details className="dropdown dropdown-end">
      <summary tabIndex={0} role="button" className="btn btn-ghost btn-circle avatar">
        <span className="sr-only">Option Menu</span>
        <UserIcon src={profileImage} />
      </summary>
      <ul className="mt-3 z-[1] p-2 shadow menu menu-sm dropdown-content bg-base-100 rounded-box w-52">
        <li>
          <button type="button" onClick={redirectProfile} className="w-full text-left">
            Perfil
          </button>
        </li>
        <li>
          <button type="button" className="w-full text-left">
            Editar Perfil
          </button>
        </li>
        <li>
          <button type="button" onClick={handleLogout} className="w-full text-left">
            Logout
          </button>
        </li>
      </ul>
    </details>
  );
}
