import React from "react";
import { API } from "../../config/Api";
import { FormatNumber } from "../../helper/Index";

const Table = ({dataCar, handleUpdate, getData}) => {

  const deleteCar = async id => {
    try {
        const response = await API.delete(`/car/${id}`)
        if(response){
            alert('car deleted successfully')
        }
        await getData();
    }catch (err) {
        console.log(JSON.stringify(err, null, 2));
    }
  }
  return (
    <>
      <div className="mt-4">
        <table className="table-cell w-full">
          <thead>
            <tr>
              <th className="px-4 py-2">Car name</th>
              <th className="px-4 py-2">Price/day</th>
              <th className="px-4 py-2">Price/month</th>
              <th className="px-4 py-2">Image</th>
              <th className="px-4 py-2">Option</th>
            </tr>
          </thead>
          <tbody>
            {dataCar.map((item) => (
              <tr key={item.car_id}>
                <td className="border px-4 py-2">{item.car_name}</td>
                <td className="border px-4 py-2">{FormatNumber(item.day_rate)}</td>
                <td className="border px-4 py-2">{FormatNumber(item.month_rate)}</td>
                <td className="border px-4 py-2">{item.image.length > 20 ? item.image.slice(0, 22) + "..." : item.image}</td>
                <td className="border px-4 py-2">
                    <button 
                    className="me-4 bg-green-500 text-white px-4 rounded-lg"
                    onClick={() => handleUpdate(item)}
                    >
                        Edit
                    </button>
                    <button 
                    className="me-4 bg-red-500 text-white px-4 rounded-lg"
                    onClick={() => deleteCar(item.car_id)}
                    >
                        Delete
                    </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
};

export default Table;