from exception.CustomException import CustomException


def try_except_logger(try_cb, finally_cb):
    try:
        try_cb()
    except CustomException as e:
        print(e)
    finally:
        finally_cb()