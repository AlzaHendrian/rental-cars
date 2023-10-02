import React, { useState } from 'react';

const InputImage = ({ onChange }) => {

  return (
    <div className="my-4">
      <div className="flex items-center space-x-2">
        <input
          type="file"
          className="border rounded p-2"
          onChange={onChange}
          name={'image'}
        />
      </div>
    </div>
  );
};

export default InputImage;
