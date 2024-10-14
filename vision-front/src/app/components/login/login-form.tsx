/* eslint-disable react/jsx-props-no-spreading */
import React, { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useRouter } from 'next/navigation';
import Credentials from '@/app/services/dtos/credentials';
import { loginAccount } from '@/app/services/user';
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

  const onSubmit: SubmitHandler<FormData> = async (data) => {
    const credentials = new Credentials(data.email, data.password);

    setErrors(credentials.validate().join(','));
    if (errors.length > 0) {
      return;
    }

    const response = await loginAccount(credentials);

    if (!response.ok) {
      setErrors(response.data.message);
      return;
    }
    router.push('/');
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
      >
        <span className="text-white text-xl font-light">Continuar com Vision</span>
      </button>
    </form>
  );
}

export default LoginForm;
