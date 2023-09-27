import React, { useState } from 'react';

const InputImage = ({ onChange }) => {
  const [selectedFile, setSelectedFile] = useState(null);

  const handleFileChange = (event) => {
    const file = event.target.files[0];
    setSelectedFile(file);
    onChange(event);
  };

  return (
    <div className="my-4">
      <div className="flex items-center space-x-2">
        <input
          type="file"
          className="border rounded p-2"
          onChange={handleFileChange}
        />
      </div>
    </div>
  );
};

export default InputImage;
