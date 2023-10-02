import { useState } from "react";
import { Input } from "../../components/Input"
import { API } from "../../config/Api";

const FormOrder = ({ carId, close }) => {
    const [formData, setFormData] = useState({
        car_id: carId,
        pickup_date: '',
        dropoff_date: '',
        pickup_location: '',
        dropoff_location: ''
    })

    const handleOnChange = (name, value) => {
        setFormData({
            ...formData,
            [name]: value
        })
    };

    const handleCreateOrder = () => {
        createOrder()
        alert("order created")
        close()
    }

    const dataText = [
        {id: 2, placeholder: "input pickup location", value: formData.pickup_location, param: 'pickup_location'},
        {id: 3, placeholder: "input dropoff location", value: formData.dropoff_location, param: 'dropoff_location'},
    ]

    const createOrder = async() => {
        try {
            const response = await API.post("/order", 
            formData,
                {
                    headers: {
                        "Content-Type": "application/json",
                    },
                }
            )
            console.log(response)
            setFormData({
                car_id: '',
                pickup_date: '',
                dropoff_date: '',
                pickup_location: '',
                dropoff_location: ''
            })
        }catch (err) {
            console.log(JSON.stringify(err, null, 4));
        }
    }
    return (
        <div className="items-center">
            <h1 className="text-2xl mb-5">Add Order</h1>
            <div className="flex gap-2 mb-3">
                <div className="flex items-center">
                    <label className="text-gray-700">Start date:
                        <input 
                        type="date" 
                        value={formData.pickup_date}
                        onChange={(e) => handleOnChange('pickup_date', e.target.value)}
                        name="pickup_date"
                        className="px-4 py-2 border rounded-md focus:outline-none focus:ring focus:border-blue-300"
                        />
                    </label>
                </div>
                <div className="">
                    <label className="text-gray-700">End date:
                        <input 
                        type="date" 
                        value={formData.dropoff_date}
                        onChange={(e) => handleOnChange('dropoff_date', e.target.value)}
                        name="dropoff_date"
                        className="px-4 py-2 border rounded-md focus:outline-none focus:ring focus:border-blue-300"
                        />
                    </label>
                </div>
            </div>
            
            {dataText.map((item) => (
                <div key={item.id} className="mb-2">
                    <Input
                    placeholder={item.placeholder}
                    value={item.value}
                    onChange={(e) => handleOnChange(item.param, e.target.value)}
                    name={item.param}
                    />
                </div>
            ))}
            <div className="grid grid-cols-1 mt-5">
                <button 
                className="bg-green-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:shadow-outline"
                onClick={handleCreateOrder}
                >
                    Order
                </button>
            </div>
        </div>
    )
}

export default FormOrder