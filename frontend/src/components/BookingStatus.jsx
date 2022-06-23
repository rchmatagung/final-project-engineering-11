import React, { useState, useEffect } from "react";
import api from '../api/api'

export default function BookingStatus() {
    const [status, setStatus] = useState([])

    const getUser = async () => {
        try {
        await api.get('/user/booking/status')
         .then((res) => {
            setStatus(res.data.data)
            console.log(res.data.data)
        })
        } catch (error) {
          console.log(error);
        }
    };
    
    useEffect(() => {
        getUser();
    }, []);

    return (
        <div className="bg-gray-100 mx-20">
            <div>
                <h1 className="lg:text-4xl text-3xl text-center text-gray-800 font-extrabold pt-6 mx-auto">Booking Status</h1>
            </div>
            <div className="mx-auto container py-8">
                <div className="flex flex-wrap items-center justify-center">
                    {status === null ? (
                        <p>belum ada kelas yang di booking</p>
                    ):(
                        <>
                            {status.map((booking, index) => {
                                return (
                                    <div className="mx-2 w-72 lg:mb-0 mb-8" key={index}>
                                        <div className="bg-yellow-300 rounded-xl">
                                            <div className="p-4">
                                                <div className="flex items-center">
                                                    <h2 className="text-lg font-semibold">{booking.mentor_name}</h2>
                                                </div>
                                                <p className="text-xs text-black mt-2">{booking.book_id}</p>
                                                <p className="text-xs text-gray-600 mt-2">{booking.status}</p>
                                            </div>
                                        </div>
                                    </div>
                                )
                            })}
                        </>
                    )}
                </div>
            </div>
        </div>
    );
}
