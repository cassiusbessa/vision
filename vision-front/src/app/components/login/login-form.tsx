/* eslint-disable react/jsx-props-no-spreading */
import React, { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useRouter } from 'next/navigation';
import Credentials from '@/app/services/dtos/requests/credentials';
import { loginAccount } from '@/app/services/account';
import { setTokenLocalStorage, setTokenSessionStorage } from '@/app/services/token';
import { loadProfileByToken } from '@/app/services/profile';
import getResponseMessage from '@/app/services/helpers/getResponseMessage';
import DefaultInput from '../input/default-form-input';
import DefaultCheckBox from '../input/default-form-checkbox';

interface FormData {
  email: string;
  password: string;
  remember: boolean;
}

function LoginForm() {
  const { register, handleSubmit } = useForm<FormData>();
  const router = useRouter();
  const [errors, setErrors] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const onSubmit: SubmitHandler<FormData> = async (form) => {
    setIsLoading(true);

    const credentials = new Credentials(form.email, form.password);

    setErrors(credentials.validate().join(','));
    if (errors.length > 0) {
      return;
    }

    const { data, ok, status } = await loginAccount(credentials);

    if (!ok && data) {
      setErrors(getResponseMessage(status, 'Email'));
      setIsLoading(false);
      return;
    }

    if (ok && data) {
      if (form.remember) {
        setTokenLocalStorage(data.token);
      } else {
        setTokenSessionStorage(data.token);
      }
      const profile = await loadProfileByToken();
      if (profile.ok && profile.data) {
        setIsLoading(false);
        router.push('/');
      } else {
        setIsLoading(false);
        router.push('/profile-manager');
      }
    }
  };

  return (
    <form className="h-full flex flex-col gap-4" onSubmit={handleSubmit(onSubmit)}>
      <DefaultInput register={register} type="email" placeholder="Email" autoComplete="email" data="email" />
      <DefaultInput register={register} type="password" placeholder="Senha" autoComplete="current-password" data="password" />
      <DefaultCheckBox register={register} label="Lembrar de mim" data="remember" />
      {errors.length > 0 && <p className="text-red-500 text-sm">{errors}</p>}
      <button
        type="submit"
        className="btn btn-secondary bg-[#C14080] hover:scale-[1.01] rounded-3xl mt-3 w-full"
        disabled={isLoading}
      >
        <span className="text-white text-xl font-light">Continuar com Vision</span>
      </button>
    </form>
  );
}

export default LoginForm;
