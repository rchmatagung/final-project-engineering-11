import React from "react";

export default function BookingStatus() {
    return (
        <div className="bg-gray-100 mx-20">
            <div>
                <h1 className="lg:text-4xl text-3xl text-center text-gray-800 font-extrabold pt-6 mx-auto">Booking Status</h1>
            </div>
            <div className="mx-auto container py-8">
                <div className="flex flex-wrap items-center justify-center">
                    {/* map status */}
                    <div className="mx-2 w-72 lg:mb-0 mb-8">
                        <div className="bg-yellow-300 rounded-xl">
                            <div className="p-4">
                                <div className="flex items-center">
                                    <h2 className="text-lg font-semibold">Mentor Name</h2>
                                </div>
                                <p className="text-xs text-gray-600 mt-2">Status</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}
