/*************************************************************************/
/* Octatron                                                              */
/* Copyright (C) 2015 Andreas T Jonsson <mail@andreasjonsson.se>         */
/*                                                                       */
/* This program is free software: you can redistribute it and/or modify  */
/* it under the terms of the GNU General Public License as published by  */
/* the Free Software Foundation, either version 3 of the License, or     */
/* (at your option) any later version.                                   */
/*                                                                       */
/* This program is distributed in the hope that it will be useful,       */
/* but WITHOUT ANY WARRANTY; without even the implied warranty of        */
/* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the         */
/* GNU General Public License for more details.                          */
/*                                                                       */
/* You should have received a copy of the GNU General Public License     */
/* along with this program.  If not, see <http://www.gnu.org/licenses/>. */
/*************************************************************************/

package octatron

type OctreeFormat int

const (
	Mip_R8G8B8A8_Branch16 OctreeFormat = iota
	Mip_R8G8B8A8_Branch32
	Mip_R8G8B8_Branch16
	Mip_R8G8B8_Branch32
	Mip_R5G6B5_Branch16
	Mip_R5G6B5_Branch32
	Mip_Index8_Branch16
	Mip_Index8_Branch32
	Index8_Branch16
	Index8_Branch32
)