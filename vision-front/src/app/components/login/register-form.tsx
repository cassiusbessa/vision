import { useForm, SubmitHandler } from 'react-hook-form';
import React from 'react';
import Account from '@/app/services/dtos/account';
import createAccount from '@/app/services/user';
import DefaultInput from '../input/default-form-input';

interface FormData {
  name: string;
  email: string;
  password: string;
}

function RegisterForm() {
  const { register, handleSubmit } = useForm<FormData>();

  const onSubmit: SubmitHandler<FormData> = async (data) => {
    const account = new Account(data.name, data.email, data.password);

    const errors = account.validate();
    if (errors.length > 0) {
      console.error(errors);
      return;
    }

    console.log(account.toJSON());

    await createAccount(account);
  };

  return (
    <form className="h-full flex flex-col gap-4" onSubmit={handleSubmit(onSubmit)}>
      <DefaultInput register={register} type="name" placeholder="Nome" autoComplete="name" data="name" />
      <DefaultInput register={register} type="email" placeholder="Email" autoComplete="email" data="email" />
      <DefaultInput register={register} type="password" placeholder="Senha" autoComplete="password" data="password" />
      <button
        className="btn btn-secondary bg-[#C14080] hover:scale-[1.01] rounded-3xl mt-3 w-full"
        type="submit"
      >
        <span className="text-white text-xl font-light">Criar conta</span>
      </button>
    </form>
  );
}

export default RegisterForm;
