import React, { useEffect, useRef, useState } from 'react';


const NeonSection = () => {

    const [isVisible, setIsVisible] = useState(false);
    const sectionRef = useRef(null);

    useEffect(() => {
        const observer = new IntersectionObserver(
            ([entry]) => {
                setIsVisible(entry.isIntersecting);
            },
            { threshold: 0.5 }
        );

        if (sectionRef.current) {
            observer.observe(sectionRef.current);
        }

        return () => {
            if (sectionRef.current) {
                observer.unobserve(sectionRef.current);
            }
        };
    }, [sectionRef]);
    const splitTextIntoLetters = (text, speed) => {
        return text.split('').map((char, index) => {
            const className = char === ' ' ? 'letter space' : 'letter';
            return (
                <span key={index} className={`${className} ${isVisible ? 'visible' : ''}`} style={{ animationDelay: `${index * speed}s` }}>
                    {char}
                </span>
            )
        });
    };

    return (
        <section id='Nuestros Objetivos' ref={sectionRef} className="neon-section">
            <div className="container">
                <div className="row justify-content-center">
                    <div className="col-xl-10">
                        <div className="neon-caption text-center">
                            <p className="neon-text neon-text-primary">
                                {splitTextIntoLetters("¿Listo para llevar tu compromiso con la seguridad digital al siguiente nivel?", 0.03)}
                            </p>
                            <p className="neon-text neon-text-secondary">
                                {splitTextIntoLetters("Registrate y únete a nuestra", 0.04)}
                            </p>
                            <p className="neon-text neon-text-secondary">
                                {splitTextIntoLetters("comunidad de seguridad", 0.05)}
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    );
};


export default NeonSection;