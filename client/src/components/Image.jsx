import "../app.css";
import React, { useState } from "react";
// import introVideo from "../video/intro.mp4";
const Image = () => {
  const [imageSrc, setImageSrc] = useState(null);
  const [imageUploaded, setImageUploaded] = useState(false);

  const handleUpload = (event) => {
    const file = event.target.files[0];
    const reader = new FileReader();
    reader.onloadend = () => {
      setImageSrc(reader.result);
      setImageUploaded(true);
    };
    reader.readAsDataURL(file);
  };

  const handleDownload = () => {
    const link = document.createElement("a");
    link.href = imageSrc;
    link.download = "DownloadedImage";
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  };

  return (
    <>
      <div class="ml">
        <input
          type="file"
          onChange={handleUpload}
          accept="image/*"
          id="fileInput"
          style={{ display: "none" }}
        />
        <button
          className="img-button"
          onClick={() => document.getElementById("fileInput").click()}
        >
          Upload Image
        </button>
        <h1>
          Handling <br />
          Input <br />
          Image <br />
        </h1>
        <h3>100% Automatically and Free</h3>
        <video width="320" height="240" autoPlay muted controls>
          <source src="/intro.mp4" type="video/mp4" />
          Your browser does not support the video tag.
        </video>
      </div>
      {imageUploaded && (
        <div className="home">
          {/* <div className="image">
            <img className="image-img" src={imageSrc} alt="Uploaded" />
          </div> */}
          <div className="image">
            <img className="image-img" src={imageSrc} alt="Uploaded" />
            <button className="img-button-download" onClick={handleDownload}>
              Download
            </button>
          </div>
        </div>
      )}
    </>
  );
};

export default Image;
