import React, { useState } from "react";
import api from '../api/api';

function UpdateProfile({refetch}) {
    const [showModal, setShowModal] = useState(false);
    const [name, setName] = useState('');
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [phoneNumber, setPhoneNumber] = useState('');
    const [address, setAddress] = useState('');
    const id = localStorage.getItem("id");

    const editProfile = async () => {
        try {
        await api.put(`/user/update/${id}`, JSON.stringify({
            username: username, 
            name: name,
            password: password,
            address: address,
            phone: phoneNumber,
            email: email,
        }))
        .then((res) => {
            console.log(res.data)
            setShowModal(false)
            refetch();
        })
        } catch (error) {
          console.log(error);
        }
    };

    return (
        <>
        <button
            className="bg-yellow-500 text-black active:bg-gray-300 px-3 py-2 rounded-xl shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1"
            type="button"
            onClick={() => setShowModal(true)}
        >
            Edit
        </button>
        {showModal ? (
            <>
            <div className="flex justify-center items-center overflow-x-hidden overflow-y-auto fixed inset-0 z-50 outline-none focus:outline-none">
                <div className="relative w-auto my-6 mx-auto max-w-3xl">
                <div className="border-0 rounded-lg shadow-lg relative flex flex-col w-full bg-white outline-none focus:outline-none">
                    <div className="flex items-start justify-between p-5 border-b border-solid border-gray-300 rounded-t ">
                    <h3 className="text-xl font=semibold">Update Profile Data</h3>
                    <button
                        className="bg-transparent border-0 text-black float-right"
                        onClick={() => setShowModal(false)}
                    >
                        <span className="text-black opacity-7 h-6 w-6 text-xl block bg-gray-400 py-0 rounded-full">
                        x
                        </span>
                    </button>
                    </div>
                    <div className="relative p-4 flex-auto">
                    <form className="bg-gray-200 shadow-md rounded px-8 pt-4 pb-6 w-full">
                        <label className="block text-black text-sm font-bold mb-1">
                        Username
                        </label>
                        <input 
                            className="shadow appearance-none border rounded w-full py-2 px-1 text-black" 
                            type="text"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            />
                        <label className="block text-black text-sm font-bold mb-1">
                        Nama
                        </label>
                        <input 
                            className="shadow appearance-none border rounded w-full py-2 px-1 text-black" 
                            type="text"
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                            />
                        <label className="block text-black text-sm font-bold mb-1">
                        Password
                        </label>
                        <input 
                            className="shadow appearance-none border rounded w-full py-2 px-1 text-black" 
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            />
                        <label className="block text-black text-sm font-bold mb-1">
                        Address
                        </label>
                        <input 
                            className="shadow appearance-none border rounded w-full py-2 px-1 text-black" 
                            type="text"
                            value={address}
                            onChange={(e) => setAddress(e.target.value)}
                            />
                        <label className="block text-black text-sm font-bold mb-1">
                        No HP
                        </label>
                        <input 
                            className="shadow appearance-none border rounded w-full py-2 px-1 text-black" 
                            type="tel"
                            value={phoneNumber}
                            onChange={(e) => setPhoneNumber(e.target.value)}
                            />
                        <label className="block text-black text-sm font-bold mb-1">
                        Email
                        </label>
                        <input 
                            className="shadow appearance-none border rounded w-full py-2 px-1 text-black" 
                            type="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            />
                    </form>
                    </div>
                    <div className="flex items-center justify-end p-6 border-t border-solid border-blueGray-200 rounded-b">
                    <button
                        className="text-red-500 background-transparent font-bold uppercase px-6 py-2 text-sm outline-none focus:outline-none mr-1 mb-1"
                        type="button"
                        onClick={() => setShowModal(false)}
                    >
                        Close
                    </button>
                    <button
                        className="text-white bg-yellow-500 active:bg-yellow-700 font-bold uppercase text-sm px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1"
                        type="button"
                        onClick={editProfile}
                    >
                        Submit
                    </button>
                    </div>
                </div>
                </div>
            </div>
            </>
        ) : null}
        </>
    );
};

export default UpdateProfile