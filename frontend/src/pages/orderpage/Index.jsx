import { useEffect, useState } from "react";
import { API } from "../../config/Api";
import dayjs from 'dayjs'

const OrderPage = () => {
    const [dataOrder, setDataOrder] = useState([])
    const getOrder = async () => {
        try {
            const response = await API.get("/orders")
            setDataOrder(response.data.data)
            console.log("ini response data order: ", response.data.data)
        }catch(err){
            console.log(JSON.stringify(err, null, 4));
        }
    }

    useEffect(() => {
        getOrder()
    }, [])

    const deleteOrder = async id => {
        try{
            const response = await API.delete(`/order/${id}`)
            if(response){
                alert("delete order success")
            }
            await getOrder();
        }catch(err){
            alert("delete order failed")
            console.log(JSON.stringify(err, null, 4));
        }
    }

    return (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 mt-10 mx-auto w-11/12">
            {dataOrder.map((item) => (
                <div key={item.order_id} className="bg-white drop-shadow-md rounded-md overflow-hidden hover:scale-[1.05] transition-transform duration-300 ease-in-out transform group">
                    <div>
                        <div className="w-80 h-40 overflow-hidden relative">
                            <p className="text-right text-sm text-yellow-600 mt-2 absolute top-0 right-10 bg-yellow-50 z-50 px-3 rounded-lg">Pending</p>
                            <img src={item.car.image} className="w-full h-full object-cover group-hover:scale-125 transition-transform duration-300 ease-in-out transform"/>
                        </div>
                        <div className="p-2">
                            <h1 className="font-semibold text-xl mb-2">{item.car.car_name}</h1>
                            <div className="flex gap-2 text-sm text-slate-600 mb-2">
                                <p>{dayjs(item.pickup_date).format('DD MMM YYYY')}</p>
                                <p>-</p>
                                <p>{dayjs(item.dropoff_date).format('DD MMM YYYY')}</p>
                            </div>
                            <div className="flex justify-between">
                                <div>
                                    <p className="text-sm font-semibold text-slate-500">Pickup: </p>
                                    <p>{item.pickup_location}</p>
                                </div>
                                <div>
                                    <p className="text-sm font-semibold text-slate-500 text-right">Dropoff: </p>
                                    <p>{item.dropoff_location}</p>
                                    
                                </div>
                            </div>
                        </div>
                        <div className="grid grid-cols-1 w-[90%] mx-auto my-3">
                            <button 
                            className=" bg-red-500 text-white px-4 rounded-lg"
                            onClick={() => deleteOrder(item.order_id)}
                            >
                                Delete
                            </button>
                        </div>
                    </div>
                </div>
            ))}
        </div>
    )
}

export default OrderPage