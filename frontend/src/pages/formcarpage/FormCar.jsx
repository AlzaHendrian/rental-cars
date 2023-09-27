import { Button } from "../../components/Button"
import { Input } from "../../components/Input"
import InputImage from "../../components/InputImage"

const FormCar = ({formData, submitData, onChange, isUpdate, updateData}) => {
    const data = [
        {id: 1, placeholder: 'input name car', param:'car_name', formData: formData.car_name},
        {id: 2, placeholder: 'input day rate', param:'day_rate', formData: formData.day_rate},
        {id: 3, placeholder: 'input month rate', param:'month_rate', formData: formData.month_rate},
    ]

    return (
        <div className="w-[40%]">
            <div className="grid grid-cols-2 gap-3 mb-3">
                {data.map((item) => (
                    <Input
                    key={item.id}
                    placeholder={item.placeholder}
                    value={item.formData}
                    onChange={(e) => onChange(e, item.param, e.target.value)}
                    name={item.param}
                    />
                ))}
            </div>
            <InputImage 
            onChange={(e) => onChange(e, 'image', e.target.files[0])}
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