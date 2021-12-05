#pragma once
#include "CommonHeader.h"

class aud_decode
{
public:
	void init(int index = 0, int sample = 0);
	void decode_chunk(const uint8_t* audio_in, short* audio_out, int cs_chunk);
	void encode_chunk(const short* audio_in, uint8_t* audio_out, int cs_chunk);

	int index() const
	{
		return m_index;
	}
private:
	int m_index;
	int m_sample;
};

void aud_decode_ws_chunk(const uint8_t* s, char* d, int cb_s, int cb_d);
