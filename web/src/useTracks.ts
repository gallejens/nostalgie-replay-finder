import { useEffect, useRef, useState } from 'react';
import type { Track } from './types';
import { devTracks } from './constants';

const isDevMode = import.meta.env.MODE === 'development';

export const useTracks = () => {
  const wsRef = useRef<WebSocket | null>(null);
  const [tracks, setTracks] = useState<Track[]>(isDevMode ? devTracks : []);

  useEffect(() => {
    if (wsRef.current !== null) return;

    wsRef.current = new WebSocket(
      `${isDevMode ? 'ws' : 'wss'}://${window.location.host}/ws`
    );
    wsRef.current.onmessage = event => {
      const track: Track = JSON.parse(event.data);
      setTracks(prev => [track, ...prev]);
    };

    // stop cloudflare tunnel timeout
    setInterval(() => {
      wsRef.current?.send(
        JSON.stringify({
          keepAlive: true,
        })
      );
    }, 10000);

    return () => {
      wsRef.current?.close();
    };
  }, []);

  useEffect(() => {
    if (tracks.length > 0) return;

    fetch(`${window.location.origin}/initial`)
      .then(response => response.json())
      .then((data: Track[]) => {
        setTracks(data.reverse());
      });
  });

  return tracks;
};
