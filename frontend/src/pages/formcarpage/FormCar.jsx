import { useState } from "react"
import { Button } from "../../components/Button"
import { Input } from "../../components/Input"
import InputImage from "../../components/InputImage"

const FormCar = ({formData, submitData, onChange, isUpdate, updateData}) => {
    const FormatPrice  = (value) => {
        if (isNaN(value)) {
            return value;
          }
        return new Intl.NumberFormat("id-ID", {
            style: 'currency',
            currency: 'IDR'
          }).format(value);
    }
    const data = [
        {id: 1, placeholder: 'input name car', param:'car_name', formData: formData.car_name},
        {id: 2, placeholder: 'input day rate', param:'day_rate', formData: formData.day_rate},
        {id: 3, placeholder: 'input month rate', param:'month_rate', formData: formData.month_rate},
    ]

    console.log(formData, "<<< ini data di formcar")
    console.log(FormatPrice(data[2].formData));

    return (
        <div className="w-[40%]">
            <div className="grid grid-cols-2 gap-3 mb-3">
                {/* {data.map((item) => (
                    <Input
                    key={item.id}
                    placeholder={item.placeholder}
                    value={FormatPrice(data[2].formData)}
                    onChange={onChange}
                    name={item.param}
                    />
                ))} */}
                <Input
                    placeholder='test'
                    value={FormatPrice(formData.car_name)}
                    onChange={onChange}
                    name='car_name'
                />
                <Input
                    placeholder='test2'
                    value={FormatPrice(formData.day_rate)}
                    onChange={onChange}
                    name='day_rate'
                />
                <Input
                    placeholder='test3'
                    value={FormatPrice(formData.month_rate)}
                    onChange={onChange}
                    name='month_rate'
                />
            </div>
            <InputImage 
            onChange={onChange}
            />
            <div className="grid grid-cols-1">
                <Button
                isUpdate={isUpdate}
                onSubmit={isUpdate ? updateData : submitData}
                />
            </div>
        </div>
    )
}

export default FormCar