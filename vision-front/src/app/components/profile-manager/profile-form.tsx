/* eslint-disable react/no-array-index-key */
/* eslint-disable react/jsx-props-no-spreading */

'use client';

import React, { useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';
import { FrameWork } from '@/app/interfaces';
import { loadTags } from '@/app/services/tag';
import Profile from '@/app/services/dtos/requests/profile';
import { createProfile } from '@/app/services/profile';
import { useRouter } from 'next/navigation';
import getResponseMessage from '@/app/services/helpers/getResponseMessage';
import DefaultInput from '../input/default-form-input';
import FrameworksDropdown from '../input/frameworks-dropdown';

function ProfileForm() {
  const router = useRouter();
  const { register, handleSubmit } = useForm();
  const [selectedFrameworks, setSelectedFrameworks] = useState<string[]>();
  const [frameworks, setFrameworks] = useState<FrameWork[]>([]);
  const [errors, setErrors] = useState<string>('');

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

  const onSubmit = async (data:any) => {
    const profile = new Profile(
      data.title,
      data.name,
      data.image,
      data.description,
      selectedFrameworks || [],
      data.link,
    );

    setErrors(profile.validate().join(','));
    if (errors.length > 0) {
      return;
    }

    const response = await createProfile(profile);
    console.log(response);

    if (!response.ok) {
      setErrors(getResponseMessage(response.status, 'Perfil'));
      return;
    }

    if (response.ok) {
      router.push('/');
    }
  };

  return (
    <form className="h-full gap-4 flex flex-col" onSubmit={handleSubmit(onSubmit)}>
      <DefaultInput register={register} type="text" placeholder="Nome Exibido" autoComplete="name" data="name" />
      <DefaultInput register={register} type="text" placeholder="Título Profissional" autoComplete="title" data="title" />
      <DefaultInput register={register} type="text" placeholder="Link para o Perfil" autoComplete="profile" data="link" />
      <DefaultInput register={register} type="text" placeholder="Link da Imagem" autoComplete="image" data="image" />
      <textarea className="w-full border-2 rounded-3xl p-4 bg-[#4f4f4f] placeholder-white" placeholder="Descrição" {...register('description', { required: true })} />
      {frameworks.length > 0
            && (
            <FrameworksDropdown
              frameWorks={frameworks}
              setSelectedFrameworks={setSelectedFrameworks}
            />
            )}

      <button
        type="submit"
        className="btn btn-secondary bg-[#C14080] hover:scale-[1.01] rounded-3xl mt-3 w-full"
        onSubmit={handleSubmit(onSubmit)}
      >
        <span className="text-white text-xl font-light">Confirmar Atualizações</span>
      </button>
    </form>
  );
}

export default ProfileForm;
