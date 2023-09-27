const CarPage = ({ dataCar }) => {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 mt-10 mx-auto w-11/12">
      {dataCar.map((item) => (
        <div key={item.car_id} className="bg-white rounded-lg shadow-md">
          <img src={item.image} alt={item.car_name} className="w-full h-40 object-cover rounded-t-lg" />
          <div className="p-4">
            <h1 className="text-xl font-semibold mb-2">{item.car_name}</h1>
            <div className="grid grid-rows-2">
              <p className="text-lg">{item.day_rate}/day</p>
              <p className="text-lg">{item.month_rate}/month</p>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}

export default CarPage;
