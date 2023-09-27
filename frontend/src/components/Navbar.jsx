import React, { useContext } from 'react';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { UserContext } from '../context/UserContext';

const data = [
    {id: 1, title: 'car page', path: '/show-car'},
    {id: 2, title: 'add car', path: '/add-car'},
]

const Navbar = () => {
    const [_, userDispatch] = useContext(UserContext);
    const location = useLocation();
    const navigate = useNavigate()

    const handleLogout = () => {
        userDispatch({ type: 'LOGOUT' });
        navigate("/login");
      };
    return (
        <>
        <div className='flex justify-between w-[90%] mx-auto mt-5'>
            <h1 className='text-3xl font-semibold'>Hellow Car</h1>
            <div className='flex gap-12 items-center'>
                <button 
                 className="bg-red-500 text-white font-semibold py-2 px-4 rounded-md"
                onClick={handleLogout}
                >
                    Logout
                </button>
            </div>
        </div>
        <div className="grid grid-cols-2 w-[50vh] place-items-center mx-auto">
            {data.map((item) => (
                <div key={item.id}>
                    <Link 
                    to={item.path}
                    className={location.pathname === item.path ? 'border-b-2 px-5 py-2 border-blue-500' : ''}
                    >
                        {item.title}
                    </Link>
                </div>
            ))}
        </div>
        </>
    )
}

export default Navbar