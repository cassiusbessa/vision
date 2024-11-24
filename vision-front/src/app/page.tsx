'use client';

import React from 'react';
import Header from './components/header';
import { projectsMock } from './mocks';
import { ProjectPostContainer } from './components/project';
import NewPostInput from './components/input/new-post-input';
import CreateProjectModal from './components/modal/create-project-modal';
import withProfile from './components/hoc/withProfile';

function Home() {
  return (
    <div className="bg-base-200 flex flex-col items-center
    min-h-screen h-screen overflow-y-scroll"
    >
      <Header />
      <NewPostInput className="mt-6" />
      <CreateProjectModal />
      <ProjectPostContainer projectsInfo={projectsMock} />
    </div>
  );
}

export default withProfile(Home);
