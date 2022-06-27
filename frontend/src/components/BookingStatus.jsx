import React, { useState, useEffect } from "react";
import api from '../api/api'

export default function BookingStatus() {
    const [status, setStatus] = useState([])
    const [showModal, setShowModal] = useState(false);
    const [mentorContact , setMentorContact ] = useState(null)

    const handleMentorContact = async (bookId) => {
        try {
            await api.get(`/user/booking/mentor/kontak/${bookId}`)
             .then((res) => {
                setMentorContact(res.data.data)
                console.log(res.data.data)
                setShowModal(true)
            })
        } catch (error) {
            console.log(error);
        }
      }

    const getStatus = async () => {
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
        getStatus();
    }, []);

    return (
        <div className="bg-gray-100">
            <div>
                <h1 className="lg:text-4xl text-3xl text-center text-gray-800 font-extrabold pt-6 mx-auto">Booking Status</h1>
            </div>
            <div className="mx-auto container py-8">
                <div className="flex flex-wrap items-center justify-center">
                    {status === null ? (
                        <p>belum ada kelas yang di booking</p>
                    ):(
                        <>
                            {status.map((booking) => {
                                return (
                                    <div className="mx-2 w-72 mb-8" key={booking.book_id}>
                                        <div className="bg-yellow-300 rounded-xl">
                                            <div className="p-4">
                                                <div className="flex items-center">
                                                    <h2 className="text-lg font-semibold">{booking.mentor_name}</h2>
                                                </div>
                                                <p className="text-xs text-black mt-2">{booking.book_id}</p>
                                                <p className="text-xs text-gray-600 mt-2">{booking.status}</p>
                                            </div>
                                            {booking.status === "Accepted" ? (
                                                <>
                                                <button
                                                className="bg-yellow-500 hover:bg-yellow-300 text-white active:bg-yellow-400 font-bold uppercase text-sm px-6 py-3 rounded-lg shadow hover:shadow-lg outline-none focus:outline-none mx-2 mb-2 ease-linear transition-all duration-150"
                                                onClick={() => handleMentorContact(booking.book_id)}>
                                                Mentor Contact
                                                </button>
                                                {showModal  ? (
                                                <>
                                                <div className="justify-center items-center flex overflow-x-hidden overflow-y-auto fixed inset-0 z-50 outline-none focus:outline-none">
                                                    <div className="relative w-auto my-6 mx-auto max-w-3xl">
                                                        <div className="border-0 rounded-lg shadow-lg w-full relative flex flex-col bg-white outline-none focus:outline-none">
                                                            <div className="relative p-6 flex-auto">
                                                                <p className="my-4 text-slate-500 text-lg leading-relaxed"> {mentorContact.name}</p>
                                                                <p className="my-4 text-slate-500 text-lg leading-relaxed"> {mentorContact.phone}</p>
                                                                <p className="my-4 text-slate-500 text-lg leading-relaxed"> {mentorContact.email}</p>
                                                                <p className="my-4 text-slate-500 text-lg leading-relaxed"> {mentorContact.address}</p>
                                                            </div>
                                                            <div className="flex items-center justify-end p-6 border-t border-solid border-slate-200 rounded-b">
                                                                <button
                                                                    className="bg-yellow-500 text-white active:bg-emerald-600 font-bold uppercase text-sm px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150"
                                                                    type="button"
                                                                    onClick={() => setShowModal(false)}
                                                                >
                                                                Close
                                                                </button>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </div>
                                                <div className="opacity-25 fixed inset-0 z-40 bg-black"></div>
                                                </>
                                                ) : null}
                                                </>
                                            ) : null}
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
