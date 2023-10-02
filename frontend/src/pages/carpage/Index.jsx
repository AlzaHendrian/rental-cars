import { useState } from "react";
import { ButtonSubmit } from "../../components/ButtonSubmit";
import FormOrder from "./FormOrderCar";
import { FormatNumber } from "../../helper/Index";

const CarPage = ({ dataCar }) => {
  const [isFormOrderVisible, setIsFormOrderVisible] = useState(false);
  const [selectedCarId, setSelectedCarId] = useState(null);

  const openFormOrder = (carId) => {
    setSelectedCarId(carId);
    setIsFormOrderVisible(true);
  }

  const closeFormOrder = () => {
    setSelectedCarId(null); 
    setIsFormOrderVisible(false);
  }


  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 mt-10 mx-auto w-11/12">
      {dataCar.map((item) => (
        <div key={item.car_id} className="bg-white drop-shadow-md rounded-md overflow-hidden transition-transform duration-300 ease-in-out transform group">
          <div className="w-80 h-40 overflow-hidden relative">
            <img src={item.image} alt={item.car_name} className="w-full h-full object-cover group-hover:scale-125 transition-transform duration-300 ease-in-out transform" />
          </div>
          <div className="p-4">
            <h1 className="text-xl font-semibold mb-2">{item.car_name}</h1>
            <div className="flex justify-between">
              <div>
                <p className="text-sm font-semibold text-slate-500">Day :</p>
                <p className="text-lg">Rp. {FormatNumber(item.day_rate)}</p>
              </div>
              <div>
                <p className="text-sm font-semibold text-slate-500 text-right">Month :</p>
                <p className="text-lg">Rp. {FormatNumber(item.month_rate)}</p>
              </div>
            </div>
          </div>
          <div className="grid grid-cols-1 w-[95%] mx-auto mb-3">
            <ButtonSubmit 
            title="Checkout"
            onClick={() => openFormOrder(item.car_id)}
            />
          </div>
        </div>
      ))}
      {/* Menampilkan FormOrder sebagai modal jika isFormOrderVisible true */}
      {isFormOrderVisible && (
        <div className="fixed inset-0 flex items-center justify-center z-50">
          <div className="absolute inset-0 bg-black opacity-60"></div>
            <div className="bg-white w-96 p-4 rounded-lg shadow-lg z-50">
              <button className="text-black rounded text-2xl" style={{ float: 'right' }} onClick={closeFormOrder}>X</button>
              <FormOrder 
              carId={selectedCarId}
              close={closeFormOrder}/>
            </div>
        </div>
      )}
    </div>
  );
}

export default CarPage;
