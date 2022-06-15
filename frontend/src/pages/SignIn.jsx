import React, {useState} from 'react'

const SignIn = (props) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  return (
    <div className="h-screen bg-gradient-to-r from-yellow-300 to-yellow-500 pt-1 md:pt-12 pb-4 px-2 md:px-0">
        <div className="max-w-lg mx-auto">
            <h1 className="text-4xl font-bold text-white text-center">Sign In</h1>
        </div>
        <div className='bg-white max-w-lg mx-auto p-8 md:p-12 my-5 rounded-xl shadow-2xl'>
            <div className="mt-5">
                <form className="flex flex-col">
                    <div className="mb-6 pt-3 rounded bg-gray-200">
                        <label className="block text-gray-700 text-sm font-bold mb-2 ml-3" for="username">Username</label>
                        <input type="text" id="username"
                        className="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-yellow-400 transition duration-500 px-3 pb-3"
                        onChange={(e) => setUsername(e.target.value)}
                        />
                    </div>
                    <div className="mb-6 pt-3 rounded bg-gray-200">
                        <label className="block text-gray-700 text-sm font-bold mb-2 ml-3" for="password">Password</label>
                        <input type="password" id="password" 
                        className="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-yellow-400 transition duration-500 px-3 pb-3"
                        onChange={(e) => setPassword(e.target.value)}
                        />
                    </div>
                    <button className="bg-yellow-400 hover:bg-yellow-500 text-white font-bold py-2 rounded shadow-lg hover:shadow-xl transition duration-200" 
                    type="submit"
                    >Sign In</button>
                    <div className="max-w-lg mx-auto text-center mt-12 mb-5">
                        <p className="text-grey">Don't have an account? <a href="/" class="text-yellow-400 font-bold hover:underline">Sign Up</a>.</p>
                    </div>
                </form>
            </div>
        </div>
    </div>
  )
}

export default SignIn