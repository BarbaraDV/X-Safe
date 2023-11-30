import React, * as react from "react";
import { pb } from "$lib";
import "./carrusel.css";
import type { RecordModel } from "pocketbase";

const Slide = (record: RecordModel) => {
  const { index, id, title, attachments, description } = record;
  return (
    <div
      className="slide"
      style={{
        backgroundImage: `url("${
          attachments?.length ? pb.getFileUrl(record, attachments[0]) : ""
        }")`,
      }}
    >
      <div className="number">{index < 10 ? `0${index}` : index}</div>
      <div className="body">
        <div className="name">{title}</div>
        <div
          className="headline"
          dangerouslySetInnerHTML={{ __html: description }}
        ></div>
        <a href={`/vulnerabilities/${id}`}>
          <div className="link">Ver más</div>
        </a>
      </div>
    </div>
  );
};

const Carrusel = () => {
  const [currentPage, setCurrentPage] = react.useState(0);
  const [dataSlides, setDataSlides] = react.useState<RecordModel[]>([]);

  react.useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await pb.collection("vulnerabilities").getList(1, 10, {
          filter:
            '(severity = "critic" || severity = "high") && deleted = false',
          sort: "created",
        });
        setDataSlides(data.items);
      } catch (error) {
        console.error("Error al cargar datos:", error);
      }
    };

    fetchData();
  }, []);

  const slideCount = dataSlides.length;
  const slidesPerPage = () =>
    window.innerWidth > 1700 ? 4 : window.innerWidth > 1200 ? 3 : 2;
  const maxPageCount = () => Math.floor(slideCount / slidesPerPage());

  const goToPage = (pageNumber = 0) => {
    setCurrentPage(Math.min(maxPageCount(), Math.max(0, pageNumber)));
  };

  react.useEffect(() => {
    const nextHandler = () => goToPage(currentPage + 1);
    const prevHandler = () => goToPage(currentPage - 1);

    const nextButton = document.getElementById("next");
    const prevButton = document.getElementById("prev");

    if (nextButton && prevButton) {
      nextButton.addEventListener("click", nextHandler);
      prevButton.addEventListener("click", prevHandler);
    }

    return () => {
      if (nextButton && prevButton) {
        nextButton.removeEventListener("click", nextHandler);
        prevButton.removeEventListener("click", prevHandler);
      }
    };
  }, [currentPage, goToPage]);

  const carruselRef = react.useRef<HTMLDivElement | null>(null);

  react.useEffect(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        const entry = entries[0];
        if (entry.isIntersecting && carruselRef.current) {
          carruselRef.current.classList.add("visible");
        }
      },
      {
        rootMargin: "0px",
        threshold: 0.1,
      }
    );

    if (carruselRef.current) {
      observer.observe(carruselRef.current);
    }

    return () => {
      if (carruselRef.current) {
        observer.unobserve(carruselRef.current);
      }
    };
  }, []);

  return (
    <div id="hero-slides" style={{ "--page": currentPage }} ref={carruselRef}>
      <div id="header">
        <div
          className="heading-text neon-text"
          style={{ padding: "0px 1.5vw 0px 0px" }}
        >
          Top
        </div>
        <div
          className="heading-text x-safe-text"
          style={{ padding: "0px 1.5vw 0px 0px" }}
        >
          10
        </div>
        <div className="heading-text neon-text">Vulnerabilidades</div>
      </div>
      <div id="slides-cont">
        <div className="button" id="next"></div>
        <div className="button" id="prev"></div>
        <div id="slides">
          {dataSlides.map((item, index) => (
            <Slide key={item.id} index={index + 1} {...item} />
          ))}
          <div id="next-catch"></div>
          <div id="prev-catch"></div>
        </div>
      </div>
    </div>
  );
};

export default Carrusel;
