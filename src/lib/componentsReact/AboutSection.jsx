import React, { useEffect, useRef, useState } from 'react';
import './sections.css'

const AboutSection = ({ title, subtitle, email, emailText, imageSrc }) => {
    const sectionRef = useRef(null);
    const [isVisible, setIsVisible] = useState(false);

    useEffect(() => {
        const observer = new IntersectionObserver(
            (entries) => {
                const [entry] = entries;
                if (entry.isIntersecting) {
                    setIsVisible(true);
                    observer.unobserve(sectionRef.current);
                }
            },
            { threshold: 0.1 }
        );

        if (sectionRef.current) {
            observer.observe(sectionRef.current);
        }

        return () => {
            if (sectionRef.current) {
                observer.unobserve(sectionRef.current);
            }
        };
    }, []);
    return (
        <section id='Acerca de nosotros' className={`objectives-section intro-section ${isVisible ? 'visible' : ''}`} ref={sectionRef}>
            <div className="container">
                <div className="row clearfix">
                    {/* Content Column */}

                    <div className="content-inner">
                        <div className="heading-style">
                            <div className="heading-text neon-text">{title}</div>
                            <h2 className='x-safe-text subtitle-text'>{subtitle}</h2>
                        </div>
                        <div className="neon-text description-text">{`El enfoque principal de este proyecto es proporcionar una comprensión detallada y actualizada de los incidentes de seguridad en la plataforma X. Nuestro enfoque se centra en apoyar la investigación académica y el conocimiento público, ofreciendo un repositorio completo y accesible para todo tipo de usuarios.`}</div>
                        <div className="neon-text description-text">{`"El único camino seguro es asumir que cada cosa es insegura hasta que se demuestre lo contrario." Bruce Schneier `}</div>
                        <div className="neon-text contact-email">
                            {emailText}: <span className="x-safe-text contact-email">{email}</span>
                        </div>

                    </div>

                    {/* Image Column */}
                    <div className="image-col">
                        <div className="image-inner">
                            <div className="image-frame">
                                <img className="frame-image" src={imageSrc} alt="" />
                            </div>
                        </div>
                    </div>

                </div>
            </div>
        </section>
    );
};

export default AboutSection;
