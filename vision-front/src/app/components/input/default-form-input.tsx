/* eslint-disable react/jsx-props-no-spreading */
/* eslint-disable react/require-default-props */
import React from 'react';
import { UseFormRegister, Path, FieldValues } from 'react-hook-form';

interface DefaultInputProps<T extends FieldValues> {
  register: UseFormRegister<T>;
  type: string;
  placeholder: string;
  autoComplete: string;
  data: Path<T>;
  className?: string;
}

function DefaultInput<T extends FieldValues>({
  register, type, placeholder, autoComplete, data, className,
}: DefaultInputProps<T>) {
  return (
    <input
      className={`w-full border-2 rounded-3xl p-4 bg-[#4f4f4f] placeholder-white ${className}`}
      type={type}
      placeholder={placeholder}
      autoComplete={autoComplete}
      {...register(data, { required: true })}
    />
  );
}

export default DefaultInput;
