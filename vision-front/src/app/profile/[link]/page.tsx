'use client';

import React, { useEffect, useState } from 'react';
import { Josefin_Sans } from 'next/font/google';
import { useRouter } from 'next/navigation';
import { loadProfileByUrl } from '@/app/services/profile';
import { LoadedProfile } from '@/app/services/dtos/responses/default-response';
import ProfileCard from '../../components/profile-card';
import { ProjectProfileContainer, StarProject } from '../../components/project';
import { projectsMock } from '../../mocks';

const inter = Josefin_Sans({ subsets: ['latin'] });

export default function Profile({ params: { link } }: { params: { link: string } }) {
  const [profile, setProfile] = useState<LoadedProfile | null>(null);
  const router = useRouter();

  useEffect(() => {
    async function fetchData() {
      const profileResponse = await loadProfileByUrl(link);
      if (profileResponse.ok && profileResponse.data) {
        setProfile(profileResponse.data);
        return;
      }

      if (!profileResponse.ok) {
        router.push('/404');
      }
    }
    fetchData();
  }, [link, router]);

  return (
    <div className={`${inter.className} bg-base-200 flex flex-col items-center min-h-screen`}>
      {/* <Header /> */}
      { profile
        && (
        <>
          <ProfileCard
            className="mt-8 md:h-[32rem] w-11/12 h-fit bg-base-100"
            userFullName={profile.profile.name}
            userTitle={profile.profile.title}
            userImage={profile.profile.image}
            bio={profile.profile.description}
          />
          <StarProject project={projectsMock[0].project} className="w-11/12 my-12 h-fit md:h-[450px] bg-base-300" />
          <ProjectProfileContainer projectsInfo={projectsMock} className="w-11/12 my-12 bg-base-100" />
        </>
        )}
    </div>
  );
}
