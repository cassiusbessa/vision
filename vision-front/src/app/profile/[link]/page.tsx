'use client';

import React, { useEffect, useState } from 'react';
import { Josefin_Sans } from 'next/font/google';
import { useRouter } from 'next/navigation';
import { loadProfileByUrl } from '@/app/services/profile';
import { LoadedProfile, ProjectDTO } from '@/app/services/dtos/responses/default-response';
import Header from '@/app/components/header';
import { useAuth } from '@/app/state/auth-context';
import { loadProjectsByProfileId } from '@/app/services/projects';
import ProfileCard from '../../components/profile-card';
import { ProjectProfileContainer, StarProject } from '../../components/project';

const inter = Josefin_Sans({ subsets: ['latin'] });

export default function Profile({ params: { link } }: { params: { link: string } }) {
  const [profile, setProfile] = useState<LoadedProfile | null>(null);
  const [projects, setProjects] = useState<ProjectDTO[] | null>([]);
  const [startProject, setStartProject] = useState<ProjectDTO | null>(null);
  const { me } = useAuth();
  const [myself, setMyself] = useState<boolean>(false);
  const router = useRouter();

  useEffect(() => {
    async function fetchData() {
      const profileResponse = await loadProfileByUrl(link);
      if (profileResponse.ok && profileResponse.data) {
        setProfile(profileResponse.data);
        if (me && me.profile.link === link) {
          setMyself(true);
        }
        const projectsResponse = await loadProjectsByProfileId(profileResponse.data.profile.id);
        console.log(projectsResponse);
        if (projectsResponse.ok && projectsResponse.data) {
          setProjects(projectsResponse.data.projects);
          const startProjectFinded = projectsResponse.data.projects.find(
            (project) => project.id === profileResponse.data?.profile.startProjects,
          ) || null;
          setStartProject(startProjectFinded);
        }
        return;
      }

      if (!profileResponse.ok) {
        router.push('/404');
      }
    }
    fetchData();
  }, [link, me, router]);

  return (
    <div className={`${inter.className} bg-base-200 flex flex-col items-center min-h-screen`}>
      <Header />
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
          {startProject && (
          <StarProject
            project={startProject}
            className="w-11/12 my-12 bg-base-100"
          />
          )}
          {projects && projects.length > 0 && (
            <ProjectProfileContainer projects={projects} className="w-11/12 my-12 bg-base-100" />
          )}
        </>
        )}
    </div>
  );
}
