'use client';

/* eslint-disable react/jsx-props-no-spreading */
import React, { useState, useEffect } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { FrameWork } from '@/app/interfaces';
import { loadTags } from '@/app/services/tag';
import Project from '@/app/services/dtos/requests/project';
import { useAuth } from '@/app/state/auth-context';
import { createProject } from '@/app/services/projects';
import UserIcon from '../user-icon';
import DefaultInput from '../input/default-form-input';
import FrameworksDropdown from '../input/frameworks-dropdown';

interface FormData {
  title: string;
  description: string;
  image: string;
  github: string;
  demo: string;
}

function CreateProjectModal() {
  const { register, handleSubmit, reset } = useForm<FormData>();
  const [frameworks, setFrameworks] = useState<FrameWork[]>([]);
  const [selectedFrameworks, setSelectedFrameworks] = useState<string[]>();
  const [errors, setErrors] = useState<string[]>([]);
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const { me } = useAuth();

  useEffect(
    () => {
      const fetchFrameworks = async () => {
        const response = await loadTags();
        const data = response.data as FrameWork[];
        setFrameworks(data);
      };
      fetchFrameworks();
    },
    [],
  );

  const profileImage = me?.profile?.image || 'https://e7.pngegg.com/pngimages/84/165/png-clipart-united-states-avatar-organization-information-user-avatar-service-computer-wallpaper-thumbnail.png';

  const handleCloseModal = () => {
    const element = document.getElementById('create_projet_modal') as any;
    if (element) element.close();
    reset();
    setSelectedFrameworks([]);
    setErrors([]);
    setIsSubmitting(false);
  };

  const onSubmit: SubmitHandler<FormData> = async (data: FormData) => {
    setIsSubmitting(true);

    const project = new Project(
      me?.account.accountId || '',
      data.title,
      data.description,
      data.github,
      data.demo,
      data.image,
      selectedFrameworks || [],
    );

    const projectErrors = project.validate();

    if (projectErrors.length > 0) {
      setErrors(projectErrors);
      return;
    }

    const response = await createProject(project);
    console.log(response);

    handleCloseModal();
  };

  return (
    <div>
      <dialog id="create_projet_modal" className="modal">
        <div className="modal-box">
          <UserIcon src={profileImage} className="mb-8" />
          {errors.length > 0 && <p className="text-red-500 text-sm">{errors}</p>}
          <div className="modal-action flex-col">
            <form className="flex flex-col gap-4" onSubmit={handleSubmit(onSubmit)}>
              <DefaultInput register={register} type="title" placeholder="Título do projeto" autoComplete="title" data="title" className="w-full" />
              <textarea className="w-full border-2 rounded-3xl p-4 bg-[#4f4f4f] placeholder-white" placeholder="Descrição" {...register('description', { required: true })} />
              <DefaultInput register={register} type="text" placeholder="Imagem" autoComplete="image" data="image" />
              <DefaultInput register={register} type="text" placeholder="Github" autoComplete="github" data="github" />
              <DefaultInput register={register} type="text" placeholder="Demonstração" autoComplete="demo" data="demo" />
              <FrameworksDropdown
                setSelectedFrameworks={setSelectedFrameworks}
                frameWorks={frameworks}
              />
              <button
                type="submit"
                className="btn btn-secondary bg-[#C14080] hover:scale-[1.01] rounded-3xl mt-3 w-full"
                onSubmit={handleSubmit(onSubmit)}
                disabled={isSubmitting}
              >
                <span className="text-white text-xl font-light">Postar projeto</span>
              </button>
              <button type="button" className="btn btn-sm btn-circle btn-ghost absolute right-2 top-2" onClick={handleCloseModal}>✕</button>
            </form>
          </div>
        </div>
      </dialog>
    </div>
  );
}

export default CreateProjectModal;
