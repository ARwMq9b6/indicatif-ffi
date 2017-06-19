from ._indicatif import lib
from .utils import py_str_to_rust_str


class ProgressStyle:
    """ Controls the rendering style of progress bars. """

    def __init__(self):
        """ Returns the default progress bar style for bars. """
        self.inner = lib.progressstyle_default_bar()

    def set_progress_chars(self, chars: str):
        """ Sets the three progress characters `(filled, current, to do)`. """
        lib.progressstyle_set_progress_chars(self.inner,
                                             py_str_to_rust_str(chars))

    def set_template(self, tpl: str):
        """ Sets the template string for the progress bar. """
        lib.progressstyle_set_template(self.inner, py_str_to_rust_str(tpl))


class ProgressBar:
    """ A progress bar or spinner """

    def __init__(self, size: int):
        """
        Creates a new progress bar with a given length.
        This progress bar by default draws directly to stdout.
        """
        self.inner = lib.progressbar_new(size)

    def set_style(self, style: ProgressStyle):
        """ Overrides the stored style. """
        lib.progressbar_set_style(self.inner, style.inner)

    def set_position(self, pos: int):
        """ Sets the position of the progress bar. """
        lib.progressbar_set_position(self.inner, pos)

    def finish_with_message(self, msg: str):
        """ Finishes the progress bar and sets a message. """
        lib.progressbar_finish_with_message(self.inner,
                                            py_str_to_rust_str(msg))
