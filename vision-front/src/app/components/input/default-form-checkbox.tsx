/* eslint-disable react/jsx-props-no-spreading */
import React from 'react';
import { FieldValues, Path, UseFormRegister } from 'react-hook-form';

interface DefaultCheckBoxProps<T extends FieldValues> {
  register: UseFormRegister<T>;
  label: string;
  data: Path<T>;
}

function DefaultCheckBox<T extends FieldValues>({
  register, label, data,
}: DefaultCheckBoxProps<T>) {
  return (
    <label
      className="font-light text-xl hover:cursor-pointer w-full"
      htmlFor={label}
    >
      <input
        className="checkbox ml-4 border-blue-500 mr-2 bg-[#4f4f4f]"
        id={label}
        type="checkbox"
        {...register(data)}
      />
      {' '}
      {label}
    </label>
  );
}

export default DefaultCheckBox;
