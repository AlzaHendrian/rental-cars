import { useContext, useState } from "react";
import FormLogin from "./FormLogin"
import { UserContext } from "../../context/UserContext";
import { useNavigate } from "react-router-dom";
import { API, setAuthToken } from "../../config/Api";

const LoginPage = () => {
    const [_, userDispatch] = useContext(UserContext);
    let navigate = useNavigate();
    const [form, setForm] = useState({
        email: '',
        password: '',
    });

    const handleOnChange = (name, value) => {
        setForm({
            ...form,
            [name]: value
        })
      };

    const login = async () => {
        try {
            const response = await API.post('/login', form);
            console.log(response);
      
            console.log('login success', response);
      
            // send data to useContext
            userDispatch({
              type: 'LOGIN_SUCCESS',
              payload: response.data.data,
            });
      
      
            setAuthToken(response.data.data.token);
            alert("login success")
            navigate('/show-car');
            console.log(response.data.data)
      
          } catch (err) {
            alert("login failed")
            console.log(JSON.stringify(err, null, 2));
          }
    }
    return (
        <>
            <FormLogin
            login={login}
            form={form}
            handleOnChange={handleOnChange}
            />
        </>
    )
}

export default LoginPage