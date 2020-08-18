import { useState, useEffect, useRef } from "react";
import io from "socket.io-client";

const useWebSocket = (addr) => {
  useEffect(() => {
    const webSocket = io.connect(`${addr}`, { secure: false });

    webSocket.on("connect", () => {
      webSocket.emit("join-stream", window.location.href);

      webSocket.on("chat-message", () => {});

      webSocket.on("user-left", () => {});
    });
  }, [addr]);
};

const useWebRTC = () => {
  const mediaStream = useRef();
  useEffect(() => {
    navigator.mediaDevices
      .getUserMedia({
        video: true,
        audio: {
          echoCancellation: true,
        },
      })
      .then((stream) => {
        mediaStream.current.srcObject = stream;
      });
  }, []);

  return mediaStream;
};

const useBrowserAgent = () => {
  const [isSupported, setIsSupported] = useState(false);
  useEffect(() => {
    let userAgent = ((navigator && navigator.userAgent) || "").toLowerCase();
    let vendor = ((navigator && navigator.vendor) || "").toLowerCase();
    let matchChrome = /google inc/.test(vendor)
      ? userAgent.match(/(?:chrome|crios)\/(\d+)/)
      : null;
    let matchFirefox = userAgent.match(/(?:firefox|fxios)\/(\d+)/);

    setIsSupported(matchChrome !== null || matchFirefox !== null);
  }, []);

  return isSupported;
};

// // capture entire screen
// useEffect(() => {
//   navigator.mediaDevices
//     .getDisplayMedia({ video: true, audio: true })
//     .then((stream) => {
//       videoRef.current.srcObject = stream;
//     });
// }, []);

export { useWebSocket, useWebRTC, useBrowserAgent };
