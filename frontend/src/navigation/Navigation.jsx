import React, { useContext, useEffect, useState } from "react";
import Navbar from "../components/Navbar";
import CarPage from "../pages/carpage";
import LoginPage from "../pages/loginpage/Login";
import RegisterPage from "../pages/registerpage";
import { Route, Routes } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import { AddCarPage } from "../pages/formcarpage/Index";
import { UserContext } from "../context/UserContext";
import { API, setAuthToken } from "../config/Api";
import OrderPage from "../pages/orderpage/Index";

const Navigation = ({ dataCar, getData }) => {
  let navigate = useNavigate();
  const [userState, userDispatch] = useContext(UserContext);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    if (!isLoading) {
      if (userState.isLogin === false) {
        navigate('/login');
      }
    }
  }, [isLoading]);

  useEffect(() => {
    if (localStorage.token) {
      setAuthToken(localStorage.token);
      checkUser();
    } else {
      setIsLoading(false);
    }
  }, []);

  const checkUser = async () => {
    try {
      const response = await API.get('/check-auth');
      // Get user data
      let payload = response.data.data;
      // Get token from local storage
      payload.token = localStorage.token;
      // Send data to useContext
      userDispatch({
        type: 'USER_SUCCESS',
        payload,
      });
      console.log(response.data.data)
      setIsLoading(false);
    } catch (error) {
      console.log('check user failed : ', error);
      userDispatch({
        type: 'AUTH_ERROR',
      });
      setIsLoading(false);
    }
  };

  return isLoading ? null : (
    <>
      {userState.isLogin && (
        <Navbar />
      )}
      <Routes>
        <Route path="/show-car" element={<CarPage dataCar={dataCar} getData={getData} />} />
        <Route path="/add-car" element={<AddCarPage dataCar={dataCar} getData={getData} />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/show-order" element={<OrderPage />} />
      </Routes>
    </>
  );
};

export default Navigation;
