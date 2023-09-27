import { Link } from "react-router-dom"
import { ButtonSubmit } from "../../components/ButtonSubmit"
import { Input } from "../../components/Input"

const FormLogin = ({login, form, handleOnChange}) => {
    const data = [
        {id: 1, placeHolderLogin:'input your email', value: form.email, param: 'email'},
        {id: 2, placeHolderLogin:'input your password', value: form.password, param: 'password'},
    ]
    return (
        <div className="w-[70vh] shadow-xl p-5 rounded-md mx-auto mt-20">
            <h1 className="text-2xl font-semibold mb-3">Sign In</h1>
            {data.map((item) => (
                <div key={item.id} className="mb-2">
                    <Input 
                    placeholder={item.placeHolderLogin}
                    value={item.value}
                    onChange={(e) => handleOnChange(item.param, e.target.value)}
                    name={item.param}
                    />
                </div>
            ))}
            <div className="grid grid-cols-1 mt-5">
                <ButtonSubmit
                onClick={login}
                title="sign in"/>
            </div>
            <p className="text-sm text-center mt-5">
                Don't have an account? click{' '}
                <Link className="font-semibold cursor-pointer" to={'/register'}>
                    Here
                </Link>
            </p>
        </div>
    )
}
export default FormLogin