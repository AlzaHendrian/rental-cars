import { useState } from "react";
import FormCar from "./FormCar"
import Table from "./Table"
import { API } from "../../config/Api";

export const AddCarPage = ({dataCar, getData}) => {
    const [isUpdate, setIsUpdate] = useState(false)
    const [formData, setFormData] = useState({
        car_id: '',
        car_name: '',
        day_rate: '',
        month_rate: '',
        image: null
        })

        const onChange = (e, param, value) => {
            setFormData({
                ...formData,
                [param]: e.target.type === "file" ? e.target.files[0] : value,
            });
        };

          const clearFormData = () => {
            setFormData({
            car_id: '',
            car_name: '',
            day_rate: '',
            month_rate: '',
            image: null
            }, setIsUpdate(false))
        }

        const handleUpdate = (item) => {
            setFormData(item)
            setIsUpdate(true)
            console.log(item)
        }

    // post data to api
    const submitData = async() => {
        const postData = new FormData();
        postData.append('car_id', formData.car_id)
        postData.append('car_name', formData.car_name)
        postData.append('day_rate', formData.day_rate)
        postData.append('month_rate', formData.month_rate)
        postData.append('image', formData.image)

        console.log("24324324")

        try {
        const response = await API.post("/car", postData,
            {
                headers: {
                    "Content-Type": "multipart/form-data",
                },
            }
        )
        console.log(response.data.data)
        alert("success adding car")
        getData()
        clearFormData()
        }catch (err) {
        console.log(JSON.stringify(err, null, 4));
        }
    }

    // update data to api
    const updateData = async(e) => {
        e.preventDefault()
        const data = new FormData();
        data.append("car_id", formData.car_id)
        data.append("car_name", formData.car_name)
        data.append("day_rate", formData.day_rate)
        data.append("month_rate", formData.month_rate)
        data.append("image", formData.image)

        try{
            const response = await API.patch(`/car/${formData.car_id}`, data, 
                {
                    headers: {
                        "Content-Type": "multipart/form-data",
                    },
                }
            )
            console.log(response.data.data);
            alert("success updating car")
            getData()
            clearFormData()
        }catch(err){
            console.log(JSON.stringify(err, null, 4));
        }
    }
    
    return (
        <div className="mt-10 w-[95%] mx-12">
            <div className="grid grid-cols-1 place-items-center mt-12">
                <FormCar 
                dataCar={dataCar}
                formData={formData}
                onChange={onChange}
                submitData={submitData}
                isUpdate={isUpdate}
                updateData={updateData}
                />
            </div>
            <div className="grid grid-cols-1 place-items-center mt-12">
                <Table 
                dataCar={dataCar}
                handleUpdate={handleUpdate}
                getData={getData}/>
            </div>
        </div>
    )
}