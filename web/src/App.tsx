import type { FC } from 'react';
import { useTracks } from './useTracks';
import './App.css';

export const App: FC = () => {
  const tracks = useTracks();

  return (
    <div className='container'>
      <h1>Nostalgie Song History</h1>
      {(tracks.length === 0 && <p>Loading...</p>) || (
        <div className='tracklist'>
          {tracks.map(track => (
            <div
              key={track.playedAt}
              className={track.alreadyPlayed ? 'track played' : 'track'}
            >
              <p className='songname'>{track.title}</p>
              <p className='artistname'>{track.artist}</p>
              <p className='playedat'>Played on: {track.playedAt}</p>
              {track.alreadyPlayed && (
                <p className='already-played-text'>
                  This track has already been played today
                </p>
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

