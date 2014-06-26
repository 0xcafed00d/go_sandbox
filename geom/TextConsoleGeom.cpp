
#include "TextConsoleGeom.h"
#include <algorithm>


namespace textMode
{
	coord operator+ (const coord& lhs, const coord& rhs)
	{
		return coord (lhs.x + rhs.x, lhs.y + rhs.y);
	}

	coord& operator+= (coord& lhs, const coord& rhs)
	{
		lhs.x += rhs.x;
		lhs.y += rhs.y;
		return lhs;
	}

	coord operator- (const coord& lhs, const coord& rhs)
	{
		return coord (lhs.x - rhs.x, lhs.y - rhs.y);
	}

	coord& operator-= (coord& lhs, const coord& rhs)
	{
		lhs.x -= rhs.x;
		lhs.y -= rhs.y;
		return lhs;
	}

	void normalise (rectangle& rect)
	{
		if (rect.min.x > rect.max.x)
			std::swap(rect.min.x, rect.max.x);

		if (rect.min.y > rect.max.y)
			std::swap(rect.min.y, rect.max.y);
	}

	rectangle normaliseCopy (const rectangle& rect)
	{
		rectangle result = rect;
		normalise(result);
		return result;
	}

	int width (const rectangle& rect)
	{
		return rect.max.x - rect.min.x;
	}

	int height (const rectangle& rect)
	{
		return rect.max.y - rect.min.y;
	}

	coord size (const rectangle& rect)
	{
		return coord (width(rect), height(rect));
	}

	bool isNil (const rectangle& rect)
	{
		return (rect.min.x == rect.max.x) || (rect.min.y == rect.max.y);
	}

	bool isNormal (const rectangle& rect)
	{
		return (rect.min.x <= rect.max.x) && (rect.min.y <= rect.max.y);
	}

	void expand (rectangle& r, const coord& c)
	{
		r.min -= c;
		r.max += c;
	}

	rectangle expandCopy (const rectangle& r, const coord& c)
	{
		rectangle result = r;
		expand(result, c);
		return result;
	}

	void translate (rectangle& r, const coord& c)
	{
		r.min += c;
		r.max += c;
	}

	rectangle translateCopy (const rectangle& r, const coord& c)
	{
		rectangle result = r;
		translate(result, c);
		return result;
	}

	bool pointInRectangle (const rectangle& r, const coord& p)
	{
		return r.min.x <= p.x && r.min.y <= p.y && p.x < r.max.x && p.y < r.max.y;
	}

	rectangle rectangleIntersection (const rectangle& r1, const rectangle& r2)
	{
		return rectangle ( std::max (r1.min.x, r2.min.x), std::max (r1.min.y, r2.min.y), 
						   std::min (r1.max.x, r2.max.x), std::min (r1.max.y, r2.max.y));
	}

	rectangle rectangleUnion (const rectangle& r1, const rectangle& r2)
	{
		return rectangle ( std::min (r1.min.x, r2.min.x), std::min (r1.min.y, r2.min.y), 
						   std::max (r1.max.x, r2.max.x), std::max (r1.max.y, r2.max.y));
	}

	bool rectangleContains (const rectangle& rOuter, const rectangle& rInner)
	{
		return pointInRectangle (rOuter, rInner.min) && pointInRectangle (rOuter, rInner.max);
	}

	rectangle rectangleFromPosSize (const coord& pos, const coord& size)
	{
		return rectangle (pos, size + pos);
	}

	rectangle rectangleFromSize (const coord& size)
	{
		return rectangle (coord(0,0), size);
	}

	void splitRectangleV( const rectangle& r, rectangle& r1out, rectangle& r2out, int splitWidth )
	{
		int h = height(r);
		int w = width(r);
		
		if (splitWidth >= w)
		{
			r1out = r;
			r2out = rectangle();
		}
		else
		{
			r1out = rectangleFromPosSize(r.min, coord(splitWidth, h));
			r2out = rectangleFromPosSize(coord (r.min.x + splitWidth, r.min.y), coord(w - splitWidth, h));
		}
	}

	void splitRectangleH( const rectangle& r, rectangle& r1out, rectangle& r2out, int splitHeight )
	{
		int h = height(r);
		int w = width(r);

		if (splitHeight >= h)
		{
			r1out = r;
			r2out = rectangle();
		}
		else
		{
			r1out = rectangleFromPosSize(r.min, coord(w, splitHeight));
			r2out = rectangleFromPosSize(coord (r.min.x, r.min.y + splitHeight), coord(w, h - splitHeight));
		}
	}
}