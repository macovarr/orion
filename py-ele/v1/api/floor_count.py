# -*- coding: utf-8 -*-
from __future__ import absolute_import, print_function

from flask import request, g

from . import Resource
from .. import schemas


class FloorCount(Resource):

    def get(self):

        return {'count': 9573}, 200, None