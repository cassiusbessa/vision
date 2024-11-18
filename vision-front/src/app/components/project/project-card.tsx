/* eslint-disable @next/next/no-img-element */
import React from 'react';
// import projectImage from '@/app/assets/project.jpg';
import { ProjectDTO } from '@/app/services/dtos/responses/default-response';
import ProjectCardTags from './project-card-tags';
import ProjectCardLinks from './project-card-links';

export default function ProjectCard({ className, project }: { className: string,
  project: ProjectDTO }) {
  return (
    <div className={`card bg-base-100 shadow-xl ${className}`}>
      <figure className="h-3/6 overflow-hidden">
        <img src={project.image} alt="Project" className="h-full w-full object-cover" />
      </figure>
      <div className="card-body flex flex-col justify-between p-4 pb-0 h-3/6">
        <article className="prose text-ellipsis overflow-hidden ...">
          <h3 className="card-title">
            {project.title}
          </h3>
          <p className="whitespace-pre-line">{project.description}</p>
        </article>
        <div className="card-actions justify-between items-center">
          <ProjectCardLinks links={project.link} />
          <ProjectCardTags tags={project.technologies} />
        </div>
      </div>
    </div>
  );
}
