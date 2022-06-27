import React from 'react';
import logoImg from '../assets/logohi.png'

const Hero = () => {
  return (
    <div className='w-full bg-gray-100 flex flex-col justify-between'>
        <div className='max-w-[1024px] mx-auto'>
            <div className='flex flex-col justify-center md:items-start w-full px-2 my-10'>
                <h1 className='m-auto py-3 text-4xl md:text-5xl lg:text-7xl font-bold'>About<span className="text-yellow-400">Us</span></h1>
                <div className="m-auto mb-5">
                    <img src={logoImg} alt="Hicoder Logo" className="object-cover w-[300px] shadow-md" />
                </div>
                <p className='m-auto md:w-1/2 pb-5 text-center text-sm md:text-base'><b>Hicoder</b> akan membantu anda untuk mencari mentor yang berpengalaman di bidang IT. Anda akan dibimbing secara profesional dengan materi yang ter-update dan sesuai dengan kebutuhan perusahaan-perusahaan saat ini. Anda juga bisa menyampaikan pertanyaan dan melakukan konsultasi terkait kesulitan anda saat mempelajari materi atau mengerjakan project anda. Let's try with us.</p>
            </div>
        </div>
    </div>
  )
}

export default Hero