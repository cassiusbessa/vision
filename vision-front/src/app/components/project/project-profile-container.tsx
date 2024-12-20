import React from 'react';
import { ProjectDTO } from '@/app/services/dtos/responses/default-response';
import ProjectCard from './project-card';

export default function ProjectProfileContainer({ className, projects }: { className: string,
  projects: ProjectDTO[] }) {
  return (
    <div className={`rounded-2xl px-8 my-4 flex flex-col items-center ${className}`}>
      <article className="prose text-ellipsis overflow-hidden ... text-center flex justify-center p-2 h-full bg-base-200 w-full my-4">
        <h2 className="text-4xl font-bold text-center">Projetos</h2>
      </article>
      <div className="grid gap-x-4 sm:grid-cols-1 md:grid-cols-2">
        {projects.map((project) => (
          <ProjectCard key={project.id} project={project} className="my-4 md:h-[500px] h-fit bg-base-200" />
        ))}
      </div>
    </div>
  );
}
